import { PrismaClient } from "../generated/client/index.js";
import type { Player } from "../generated/client/index.d.ts";

export { Player };

export default class DB {
  private static prisma = new PrismaClient();

  public static async setPlayerColor(playerUuid: string, color: string): Promise<void> {
    await this.prisma.player.upsert({
      where: {
        playerUuid,
      },
      update: {
        color,
      },
      create: {
        playerUuid,
        color,
      },
    });
  }

  public static async getPlayerColors(): Promise<Map<string, string>> {
    const players = await this.prisma.player.findMany({
      select: {
        playerUuid: true,
        color: true,
      },
    });
    const map = new Map<string, string>();
    for (const player of players) {
      map.set(player.playerUuid, player.color);
    }
    return map;
  }

  public static async setOrCreateStat(stat: string, value: number): Promise<void> {
    await this.prisma.stats.upsert({
      where: {
        stat,
      },
      update: {
        value,
      },
      create: {
        stat,
        value,
      },
    });
  }

  public static async incOrCreateStat(stat: string, value: number): Promise<void> {
    await this.prisma.stats.upsert({
      where: {
        stat,
      },
      update: {
        value: {
          increment: value,
        },
      },
      create: {
        stat,
        value,
      },
    });
  }
}

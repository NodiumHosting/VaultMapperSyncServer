import CellType from "../vaults/CellType.ts";
import RoomName from "../vaults/RoomName.ts";
import RoomType from "../vaults/RoomType.ts";
import VaultCell from "../vaults/VaultCell.ts";

export default class CellPacket {
  public readonly x: number;
  public readonly z: number;
  public readonly i: boolean;
  public readonly c: CellType;
  public readonly r: RoomType;
  public readonly n: RoomName;
  public readonly m: boolean;

  constructor(x: number, z: number, inscripted: boolean, cellType: CellType, roomType: RoomType, roomName: RoomName, marked: boolean) {
    this.x = x;
    this.z = z;
    this.i = inscripted;
    this.c = cellType;
    this.r = roomType;
    this.n = roomName;
    this.m = marked;
  }

  public static parse(json: string): CellPacket {
    const obj = JSON.parse(json);
    return new CellPacket(obj.x, obj.z, obj.i, obj.c, obj.r, obj.n, obj.m);
  }

  public static toVaultCell(packet: CellPacket): VaultCell {
    return new VaultCell(packet.x, packet.z, packet.i, packet.c, packet.r, packet.n, packet.m);
  }

  public static fromVaultCell(cell: VaultCell): CellPacket {
    return new CellPacket(cell.x, cell.z, cell.inscripted, cell.cellType, cell.roomType, cell.roomName, cell.marked);
  }
}

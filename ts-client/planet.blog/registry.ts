import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSendIbcHorusAction } from "./types/planet/blog/tx";
import { MsgSendIbcPost } from "./types/planet/blog/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/planet.blog.MsgSendIbcHorusAction", MsgSendIbcHorusAction],
    ["/planet.blog.MsgSendIbcPost", MsgSendIbcPost],
    
];

export { msgTypes }
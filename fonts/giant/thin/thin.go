package thingiant

// Get this into a format that we can parse and generate a proper font file
//0;
// ╭━━━╮
// ┃╭━╮┃
// ┃┃ ┃┃
// ┃┃ ┃┃
// ┃╰━╯┃
// ╰━━━╯;
//1;
//  ╭╮
// ╭╯┃
// ╰╮┃
//  ┃┃
// ╭╯╰╮
// ╰━━╯;
//2;
// ╭━━━╮
// ┃╭━╮┃
// ╰╯╭╯┃
// ╭━╯╭╯
// ┃ ╰━╮
// ╰━━━╯;
//3;
// ╭━━━╮
// ┃╭━╮┃
// ╰╯╭╯┃
// ╭╮╰╮┃
// ┃╰━╯┃
// ╰━━━╯;
// ╭╮ ╭╮
// ┃┃ ┃┃
// ┃╰━╯┃
// ╰━━╮┃
//    ┃┃
//    ╰╯;
// ╭━━━╮
// ┃╭━━╯
// ┃╰━━╮
// ╰━━╮┃
// ╭━━╯┃
// ╰━━━╯;
// ╭━━━╮
// ┃╭━━╯
// ┃╰━━╮
// ┃╭━╮┃
// ┃╰━╯┃
// ╰━━━╯;
// ╭━━━╮
// ┃╭━╮┃
// ╰╯╭╯┃
//   ┃╭╯
//   ┃┃
//   ╰╯ ;
// ╭━━━╮
// ┃╭━╮┃
// ┃╰━╯┃
// ┃╭━╮┃
// ┃╰━╯┃
// ╰━━━╯;
// ╭━━━╮
// ┃╭━╮┃
// ┃╰━╯┃
// ╰━━╮┃
// ╭━━╯┃
// ╰━━━╯;
// ╭━━╮
// ┃╭╮┃
// ┃╭╮┃
// ╰╯╰╯;
// ╭╮
// ┃┃
// ┃╰━╮
// ┃╭╮┃
// ┃╰╯┃
// ╰━━╯;
// ╭━━╮
// ┃╭━╯
// ┃╰━╮
// ╰━━╯;
//   ╭╮
//   ┃┃
// ╭━╯┃
// ┃╭╮┃
// ┃╰╯┃
// ╰━━╯;
// ╭━━╮
// ┃ ━┫
// ┃ ━┫
// ╰━━╯;
//  ╭━╮
//  ┃╭╯
// ╭╯╰╮
// ╰╮╭╯
//  ┃┃
//  ╰╯ ;
// ╭━━╮
// ┃╭╮┃
// ┃╰╯┃
// ╰━╮┃
// ╭━╯┃
// ╰━━╯;
// ╭╮
// ┃┃
// ┃╰━╮
// ┃╭╮┃
// ┃┃┃┃
// ╰╯╰╯;
// ╭╮
// ┣┫
// ┃┃
// ╰╯;
//  ╭╮
//  ╰╯
//  ╭╮
//  ┃┃
//  ┃┃
// ╭╯┃
// ╰━╯;
// ╭╮
// ┃┃
// ┃┃╭╮
// ┃╰╯╯
// ┃╭╮╮
// ╰╯╰╯;
// ╭╮
// ┃┃
// ┃┃
// ┃┃
// ┃╰╮
// ╰━╯;
// ╭╮╭╮
// ┃╰╯┃
// ┃┃┃┃
// ╰┻┻╯;
// ╭━╮
// ┃╭╮╮
// ┃┃┃┃
// ╰╯╰╯;
// ╭━━╮
// ┃╭╮┃
// ┃╰╯┃
// ╰━━╯;
// ╭━━╮
// ┃╭╮┃
// ┃╰╯┃
// ┃╭━╯
// ┃┃
// ╰╯  ;
// ╭━━╮
// ┃╭╮┃
// ┃╰╯┃
// ╰━╮┃
//   ┃┃
//   ╰╯;
// ╭━╮
// ┃╭╯
// ┃┃
// ╰╯ ;
// ╭━━╮
// ┃━━┫
// ┣━━┃
// ╰━━╯;
//  ╭╮
// ╭╯╰╮
// ╰╮╭╯
//  ┃┃
//  ┃╰╮
//  ╰━╯;
// ╭╮╭╮
// ┃┃┃┃
// ┃╰╯┃
// ╰━━╯;
// ╭╮╭╮
// ┃╰╯┃
// ╰╮╭╯
//  ╰╯ ;
// ╭╮╭╮╭╮
// ┃╰╯╰╯┃
// ╰╮╭╮╭╯
//  ╰╯╰╯ ;
// ╭╮╭╮
// ╰╋╋╯
// ╭╋╋╮
// ╰╯╰╯;
// ╭╮ ╭╮
// ┃┃ ┃┃
// ┃╰━╯┃
// ╰━╮╭╯
// ╭━╯┃
// ╰━━╯ ;
// ╭━━━╮
// ┣━━ ┃
// ┃ ━━┫
// ╰━━━╯;
// ╭━━━╮
// ┃╭━╮┃
// ┃┃ ┃┃
// ┃╰━╯┃
// ┃╭━╮┃
// ╰╯ ╰╯;
// ╭━━╮
// ┃╭╮┃
// ┃╰╯╰╮
// ┃╭━╮┃
// ┃╰━╯┃
// ╰━━━╯;
// ╭━━━╮
// ┃╭━╮┃
// ┃┃ ╰╯
// ┃┃ ╭╮
// ┃╰━╯┃
// ╰━━━╯;
// ╭━━━╮
// ╰╮╭╮┃
//  ┃┃┃┃
//  ┃┃┃┃
// ╭╯╰╯┃
// ╰━━━╯;
// ╭━━━╮
// ┃╭━━╯
// ┃╰━━╮
// ┃╭━━╯
// ┃╰━━╮
// ╰━━━╯;
// ╭━━━╮
// ┃╭━━╯
// ┃╰━━╮
// ┃╭━━╯
// ┃┃
// ╰╯   ;
// ╭━━━╮
// ┃╭━╮┃
// ┃┃ ╰╯
// ┃┃╭━╮
// ┃╰┻━┃
// ╰━━━╯;
// ╭╮ ╭╮
// ┃┃ ┃┃
// ┃╰━╯┃
// ┃╭━╮┃
// ┃┃ ┃┃
// ╰╯ ╰╯;
// ╭━━╮
// ╰┫┣╯
//  ┃┃
//  ┃┃
// ╭┫┣╮
// ╰━━╯;
//   ╭╮
//   ┃┃
//   ┃┃
// ╭╮┃┃
// ┃╰╯┃
// ╰━━╯;
// ╭╮╭━╮
// ┃┃┃╭╯
// ┃╰╯╯
// ┃╭╮┃
// ┃┃┃╰╮
// ╰╯╰━╯;
// ╭╮
// ┃┃
// ┃┃
// ┃┃ ╭╮
// ┃╰━╯┃
// ╰━━━╯;
// ╭━╮╭━╮
// ┃ ╰╯ ┃
// ┃╭╮╭╮┃
// ┃┃┃┃┃┃
// ┃┃┃┃┃┃
// ╰╯╰╯╰╯;
// ╭━╮ ╭╮
// ┃ ╰╮┃┃
// ┃╭╮╰╯┃
// ┃┃╰╮ ┃
// ┃┃ ┃ ┃
// ╰╯ ╰━╯;
// ╭━━━╮
// ┃╭━╮┃
// ┃┃ ┃┃
// ┃┃ ┃┃
// ┃╰━╯┃
// ╰━━━╯;
// ╭━━━╮
// ┃╭━╮┃
// ┃╰━╯┃
// ┃╭━━╯
// ┃┃
// ╰╯   ;
// ╭━━━╮
// ┃╭━╮┃
// ┃┃ ┃┃
// ┃┃ ┃┃
// ┃╰━╯┃
// ╰━━╮┃
//    ╰╯;
// ╭━━━╮
// ┃╭━╮┃
// ┃╰━╯┃
// ┃╭╮╭╯
// ┃┃┃╰╮
// ╰╯╰━╯;
// ╭━━━╮
// ┃╭━╮┃
// ┃╰━━╮
// ╰━━╮┃
// ┃╰━╯┃
// ╰━━━╯;
// ╭━━━━╮
// ┃╭╮╭╮┃
// ╰╯┃┃╰╯
//   ┃┃
//   ┃┃
//   ╰╯  ;
// ╭╮ ╭╮
// ┃┃ ┃┃
// ┃┃ ┃┃
// ┃┃ ┃┃
// ┃╰━╯┃
// ╰━━━╯;
// ╭╮  ╭╮
// ┃╰╮╭╯┃
// ╰╮┃┃╭╯
//  ┃╰╯┃
//  ╰╮╭╯
//   ╰╯  ;
// ╭╮╭╮╭╮
// ┃┃┃┃┃┃
// ┃┃┃┃┃┃
// ┃╰╯╰╯┃
// ╰╮╭╮╭╯
//  ╰╯╰╯ ;
// ╭━╮╭━╮
// ╰╮╰╯╭╯
//  ╰╮╭╯
//  ╭╯╰╮
// ╭╯╭╮╰╮
// ╰━╯╰━╯;
// ╭╮  ╭╮
// ┃╰╮╭╯┃
// ╰╮╰╯╭╯
//  ╰╮╭╯
//   ┃┃
//   ╰╯  ;
// ╭━━━━╮
// ╰━━╮━┃
//   ╭╯╭╯
//  ╭╯╭╯
// ╭╯━╰━╮
// ╰━━━━╯;
// ╭╮
// ╰╯
// ╭╮
// ╰╯;
// ╭╮
// ╰┫
//  ╯;
// ╭╮
// ╰╯;
// ╭╮
// ┃┃
// ┃┃
// ╰╯
// ╭╮
// ╰╯;
// ╭━━━╮
// ┃╭━╮┃
// ╰╯╭╯┃
//   ┃╭╯
//   ╭╮
//   ╰╯ ;
// ╭━╮
// ╰╮╰╮
//  ╰╮╰╮
//   ┃ ┃
//   ┃ ┃
//  ╭╯╭╯
// ╭╯╭╯
// ╰━╯  ;
//   ╭━╮
//  ╭╯╭╯
// ╭╯╭╯
// ┃ ┃
// ┃ ┃
// ╰╮╰╮
//  ╰╮╰╮
//   ╰━╯;
// ╭┳╮
// ┃┃┃
// ╰┻╯
//    ;
// ╭╮
// ┃┃
// ╰╯
//   ;
// ╭━━╮
// ╰━━╯
//     ;
//  ╭╮
// ╭╯╰╮
// ╰╮╭╯
//  ╰╯ ;
// ╭━━━╮
// ╰━━━╯
// ╭━━━╮
// ╰━━━╯;

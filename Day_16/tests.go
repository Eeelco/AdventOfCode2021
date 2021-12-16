package main

import (
	"fmt"
	"os"
)

func TestPartOne() {
    pack, _ := ParsePacket(TranslateData("8A004A801A8002F478"))
    sum := SumVersions(pack)
    if sum != 16 {
        fmt.Printf("Expected 16, got %d", sum)
        os.Exit(1)
    }
    pack, _ = ParsePacket(TranslateData("620080001611562C8802118E34"))
    sum = SumVersions(pack)
    if sum != 12 {
        fmt.Printf("Expected 12, got %d", sum)
        os.Exit(1)
    }
    pack, _ = ParsePacket(TranslateData("C0015000016115A2E0802F182340"))
    sum = SumVersions(pack)
    if sum != 23 {
        fmt.Printf("Expected 23, got %d", sum)
        os.Exit(1)
    }
    pack, _ = ParsePacket(TranslateData("A0016C880162017C3686B18A3D4780"))
    sum = SumVersions(pack)
    if sum != 31 {
        fmt.Printf("Expected 31, got %d", sum)
        os.Exit(1)
    }
}

func TestPartTwo() {
    pack, _ := ParsePacket(TranslateData("C200B40A82"))
    val := EvalExpression(pack)
    if val != 3 {
        fmt.Printf("Expected 3, got %d", val)
        os.Exit(1)
    }
    pack, _ = ParsePacket(TranslateData("04005AC33890"))
    val = EvalExpression(pack)
    if val != 54 {
        fmt.Printf("Expected 54, got %d", val)
        os.Exit(1)
    }
    pack, _ = ParsePacket(TranslateData("880086C3E88112"))
    val = EvalExpression(pack)
    if val != 7 {
        fmt.Printf("Expected 7, got %d", val)
        os.Exit(1)
    }
    pack, _ = ParsePacket(TranslateData("CE00C43D881120"))
    val = EvalExpression(pack)
    if val != 9 {
        fmt.Printf("Expected 9, got %d", val)
        os.Exit(1)
    }
    pack, _ = ParsePacket(TranslateData("D8005AC2A8F0"))
    val = EvalExpression(pack)
    if val != 1 {
        fmt.Printf("LT: Expected 1, got %d", val)
        os.Exit(1)
    }
    pack, _ = ParsePacket(TranslateData("F600BC2D8F"))
    val = EvalExpression(pack)
    if val != 0 {
        fmt.Printf("GT: Expected 0, got %d", val)
        os.Exit(1)
    }
    pack, _ = ParsePacket(TranslateData("9C005AC2F8F0"))
    val = EvalExpression(pack)
    if val != 0 {
        fmt.Printf("EQ: Expected 0, got %d", val)
        os.Exit(1)
    }
    pack, _ = ParsePacket(TranslateData("9C0141080250320F1802104A08"))
    val = EvalExpression(pack)
    if val != 1 {
        fmt.Printf("Last: Expected 1, got %d", val)
        os.Exit(1)
    }
}

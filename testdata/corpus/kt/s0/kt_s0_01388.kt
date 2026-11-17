package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11420<A, B>(val first: A, val second: B)

fun combine11420(left: Int, right: Int): Pair11420<Int, Int> {
    return Pair11420(left + 10, right + 12)
}
data class Pair11421<A, B>(val first: A, val second: B)

fun combine11421(left: Int, right: Int): Pair11421<Int, Int> {
    return Pair11421(left + 17, right + 23)
}
data class Pair11422<A, B>(val first: A, val second: B)

fun combine11422(left: Int, right: Int): Pair11422<Int, Int> {
    return Pair11422(left + 24, right + 34)
}
@JvmName("step11423")
fun step11423(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
@JvmName("step11424")
fun step11424(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
@JvmName("step11425")
fun step11425(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}

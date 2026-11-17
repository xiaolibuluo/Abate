package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15130<A, B>(val first: A, val second: B)

fun combine15130(left: Int, right: Int): Pair15130<Int, Int> {
    return Pair15130(left + 80, right + 47)
}
data class Pair15131<A, B>(val first: A, val second: B)

fun combine15131(left: Int, right: Int): Pair15131<Int, Int> {
    return Pair15131(left + 87, right + 58)
}

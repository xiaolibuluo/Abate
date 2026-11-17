package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair25340<A, B>(val first: A, val second: B)

fun combine25340(left: Int, right: Int): Pair25340<Int, Int> {
    return Pair25340(left + 64, right + 57)
}
data class Pair25341<A, B>(val first: A, val second: B)

fun combine25341(left: Int, right: Int): Pair25341<Int, Int> {
    return Pair25341(left + 71, right + 68)
}

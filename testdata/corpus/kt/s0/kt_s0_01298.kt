package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10720<A, B>(val first: A, val second: B)

fun combine10720(left: Int, right: Int): Pair10720<Int, Int> {
    return Pair10720(left + 70, right + 10)
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20440<A, B>(val first: A, val second: B)

fun combine20440(left: Int, right: Int): Pair20440<Int, Int> {
    return Pair20440(left + 96, right + 43)
}

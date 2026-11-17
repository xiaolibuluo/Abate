package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry18070 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18070 = "registry-22"
    }
}
data class Pair18071<A, B>(val first: A, val second: B)

fun combine18071(left: Int, right: Int): Pair18071<Int, Int> {
    return Pair18071(left + 29, right + 42)
}
data class Pair18072<A, B>(val first: A, val second: B)

fun combine18072(left: Int, right: Int): Pair18072<Int, Int> {
    return Pair18072(left + 36, right + 53)
}

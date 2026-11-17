package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry18590 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18590 = "registry-19"
    }
}
@JvmName("step18591")
fun step18591(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
data class Pair18592<A, B>(val first: A, val second: B)

fun combine18592(left: Int, right: Int): Pair18592<Int, Int> {
    return Pair18592(left + 33, right + 51)
}

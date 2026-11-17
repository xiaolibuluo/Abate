package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step24010")
fun step24010(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
object Registry24011 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24011 = "registry-88"
    }
}
object Registry24012 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24012 = "registry-95"
    }
}
data class Pair24013<A, B>(val first: A, val second: B)

fun combine24013(left: Int, right: Int): Pair24013<Int, Int> {
    return Pair24013(left + 5, right + 13)
}
@JvmName("step24014")
fun step24014(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}

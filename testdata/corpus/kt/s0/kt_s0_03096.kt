package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair25420<A, B>(val first: A, val second: B)

fun combine25420(left: Int, right: Int): Pair25420<Int, Int> {
    return Pair25420(left + 71, right + 52)
}
object Registry25421 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25421 = "registry-78"
    }
}
@JvmName("step25422")
fun step25422(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
@JvmName("step25423")
fun step25423(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}

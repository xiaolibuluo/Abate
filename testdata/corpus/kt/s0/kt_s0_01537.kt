package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry12730 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12730 = "registry-64"
    }
}
data class Pair12731<A, B>(val first: A, val second: B)

fun combine12731(left: Int, right: Int): Pair12731<Int, Int> {
    return Pair12731(left + 71, right + 25)
}
@JvmName("step12732")
fun step12732(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}

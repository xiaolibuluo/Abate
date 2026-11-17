package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23910<A, B>(val first: A, val second: B)

fun combine23910(left: Int, right: Int): Pair23910<Int, Int> {
    return Pair23910(left + 48, right + 32)
}
object Registry23911 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23911 = "registry-55"
    }
}
@JvmName("step23912")
fun step23912(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}

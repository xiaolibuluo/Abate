package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16220")
fun step16220(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
@JvmName("step16221")
fun step16221(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 49
}
@JvmName("step16222")
fun step16222(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
object Registry16223 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16223 = "registry-63"
    }
}
data class Pair16224<A, B>(val first: A, val second: B)

fun combine16224(left: Int, right: Int): Pair16224<Int, Int> {
    return Pair16224(left + 70, right + 61)
}

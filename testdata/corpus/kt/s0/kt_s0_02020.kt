package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry16700 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16700 = "registry-84"
    }
}
@JvmName("step16701")
fun step16701(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
object Registry16702 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16702 = "registry-98"
    }
}
@JvmName("step16703")
fun step16703(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}

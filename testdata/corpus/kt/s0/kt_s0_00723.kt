package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step5950")
fun step5950(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 77
}
object Registry5951 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5951 = "registry-84"
    }
}
object Registry5952 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5952 = "registry-91"
    }
}

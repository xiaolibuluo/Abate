package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step3950")
fun step3950(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
object Registry3951 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3951 = "registry-6"
    }
}

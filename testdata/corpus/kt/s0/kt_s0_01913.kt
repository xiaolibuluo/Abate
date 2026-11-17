package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step15780")
fun step15780(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
fun <T> List<T>.secondOrNull15781(): T? = if (size > 1) this[1] else null

fun render15781(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry15782 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15782 = "registry-66"
    }
}
@JvmName("step15783")
fun step15783(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 73
}
object Registry15784 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15784 = "registry-80"
    }
}

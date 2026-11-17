package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step2150")
fun step2150(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 84
}
fun <T> List<T>.secondOrNull2151(): T? = if (size > 1) this[1] else null

fun render2151(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter2152 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
object Registry2153 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2153 = "registry-8"
    }
}

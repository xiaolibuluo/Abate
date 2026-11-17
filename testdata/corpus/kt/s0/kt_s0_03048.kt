package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25020")
fun step25020(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
object Registry25021 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25021 = "registry-43"
    }
}
fun <T> List<T>.secondOrNull25022(): T? = if (size > 1) this[1] else null

fun render25022(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry25023 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25023 = "registry-57"
    }
}
class Counter25024 {
    var current: Int by Delegates.observable(64) { _, _, new -> println(new) }

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

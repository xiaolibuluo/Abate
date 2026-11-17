package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter25760 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

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
@JvmName("step25761")
fun step25761(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
@JvmName("step25762")
fun step25762(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
fun <T, R> List<T>.mapChecked25763(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked25764(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25765(): T? = if (size > 1) this[1] else null

fun render25765(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry25766 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25766 = "registry-70"
    }
}
fun <T> List<T>.secondOrNull25767(): T? = if (size > 1) this[1] else null

fun render25767(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked25768(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

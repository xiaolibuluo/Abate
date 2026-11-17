package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry25210 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25210 = "registry-89"
    }
}
fun <T, R> List<T>.mapChecked25211(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter25212 {
    var current: Int by Delegates.observable(6) { _, _, new -> println(new) }

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
@JvmName("step25213")
fun step25213(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 13
}
fun <T> List<T>.secondOrNull25214(): T? = if (size > 1) this[1] else null

fun render25214(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry25215 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25215 = "registry-27"
    }
}
fun <T> List<T>.secondOrNull25216(): T? = if (size > 1) this[1] else null

fun render25216(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter25217 {
    var current: Int by Delegates.observable(41) { _, _, new -> println(new) }

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

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5820 {
    data class Success(val payload: String) : Outcome5820()
    data object Failure : Outcome5820()
}

fun describe5820(outcome: Outcome5820): String = when (outcome) {
    is Outcome5820.Success -> outcome.payload
    Outcome5820.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked5821(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter5822 {
    var current: Int by Delegates.observable(19) { _, _, new -> println(new) }

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
@JvmName("step5823")
fun step5823(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
class Counter5824 {
    var current: Int by Delegates.observable(33) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull5825(): T? = if (size > 1) this[1] else null

fun render5825(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry5826 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5826 = "registry-47"
    }
}
fun <T> List<T>.secondOrNull5827(): T? = if (size > 1) this[1] else null

fun render5827(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

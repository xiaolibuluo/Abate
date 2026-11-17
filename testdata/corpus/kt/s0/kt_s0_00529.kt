package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome4270 {
    data class Success(val payload: String) : Outcome4270()
    data object Failure : Outcome4270()
}

fun describe4270(outcome: Outcome4270): String = when (outcome) {
    is Outcome4270.Success -> outcome.payload
    Outcome4270.Failure -> "failed"
}
object Registry4271 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4271 = "registry-34"
    }
}
fun <T> List<T>.secondOrNull4272(): T? = if (size > 1) this[1] else null

fun render4272(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter4273 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
@JvmName("step4274")
fun step4274(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
sealed class Outcome4275 {
    data class Success(val payload: String) : Outcome4275()
    data object Failure : Outcome4275()
}

fun describe4275(outcome: Outcome4275): String = when (outcome) {
    is Outcome4275.Success -> outcome.payload
    Outcome4275.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked4276(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

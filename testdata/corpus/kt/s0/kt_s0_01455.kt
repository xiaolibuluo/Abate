package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull12020(): T? = if (size > 1) this[1] else null

fun render12020(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter12021 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
@JvmName("step12022")
fun step12022(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
sealed class Outcome12023 {
    data class Success(val payload: String) : Outcome12023()
    data object Failure : Outcome12023()
}

fun describe12023(outcome: Outcome12023): String = when (outcome) {
    is Outcome12023.Success -> outcome.payload
    Outcome12023.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12024(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12025 {
    data class Success(val payload: String) : Outcome12025()
    data object Failure : Outcome12025()
}

fun describe12025(outcome: Outcome12025): String = when (outcome) {
    is Outcome12025.Success -> outcome.payload
    Outcome12025.Failure -> "failed"
}

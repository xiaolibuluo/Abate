package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull4790(): T? = if (size > 1) this[1] else null

fun render4790(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4791(): T? = if (size > 1) this[1] else null

fun render4791(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4792(): T? = if (size > 1) this[1] else null

fun render4792(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4793(): T? = if (size > 1) this[1] else null

fun render4793(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4794(): T? = if (size > 1) this[1] else null

fun render4794(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome4795 {
    data class Success(val payload: String) : Outcome4795()
    data object Failure : Outcome4795()
}

fun describe4795(outcome: Outcome4795): String = when (outcome) {
    is Outcome4795.Success -> outcome.payload
    Outcome4795.Failure -> "failed"
}
fun <T> List<T>.secondOrNull4796(): T? = if (size > 1) this[1] else null

fun render4796(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter4797 {
    var current: Int by Delegates.observable(73) { _, _, new -> println(new) }

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

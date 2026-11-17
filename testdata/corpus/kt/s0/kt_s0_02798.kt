package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull23100(): T? = if (size > 1) this[1] else null

fun render23100(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter23101 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull23102(): T? = if (size > 1) this[1] else null

fun render23102(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome23103 {
    data class Success(val payload: String) : Outcome23103()
    data object Failure : Outcome23103()
}

fun describe23103(outcome: Outcome23103): String = when (outcome) {
    is Outcome23103.Success -> outcome.payload
    Outcome23103.Failure -> "failed"
}

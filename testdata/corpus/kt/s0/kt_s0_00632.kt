package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter5160 {
    var current: Int by Delegates.observable(20) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull5161(): T? = if (size > 1) this[1] else null

fun render5161(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull5162(): T? = if (size > 1) this[1] else null

fun render5162(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull5163(): T? = if (size > 1) this[1] else null

fun render5163(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step5164")
fun step5164(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
sealed class Outcome5165 {
    data class Success(val payload: String) : Outcome5165()
    data object Failure : Outcome5165()
}

fun describe5165(outcome: Outcome5165): String = when (outcome) {
    is Outcome5165.Success -> outcome.payload
    Outcome5165.Failure -> "failed"
}

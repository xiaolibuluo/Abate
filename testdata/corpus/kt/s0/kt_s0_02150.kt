package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull17750(): T? = if (size > 1) this[1] else null

fun render17750(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter17751 {
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
@JvmName("step17752")
fun step17752(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
sealed class Outcome17753 {
    data class Success(val payload: String) : Outcome17753()
    data object Failure : Outcome17753()
}

fun describe17753(outcome: Outcome17753): String = when (outcome) {
    is Outcome17753.Success -> outcome.payload
    Outcome17753.Failure -> "failed"
}

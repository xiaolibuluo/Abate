package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6030 {
    data class Success(val payload: String) : Outcome6030()
    data object Failure : Outcome6030()
}

fun describe6030(outcome: Outcome6030): String = when (outcome) {
    is Outcome6030.Success -> outcome.payload
    Outcome6030.Failure -> "failed"
}
@JvmName("step6031")
fun step6031(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
fun <T> List<T>.secondOrNull6032(): T? = if (size > 1) this[1] else null

fun render6032(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter6033 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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

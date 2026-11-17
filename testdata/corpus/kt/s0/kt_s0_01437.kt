package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull11840(): T? = if (size > 1) this[1] else null

fun render11840(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome11841 {
    data class Success(val payload: String) : Outcome11841()
    data object Failure : Outcome11841()
}

fun describe11841(outcome: Outcome11841): String = when (outcome) {
    is Outcome11841.Success -> outcome.payload
    Outcome11841.Failure -> "failed"
}
@JvmName("step11842")
fun step11842(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
class Counter11843 {
    var current: Int by Delegates.observable(92) { _, _, new -> println(new) }

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

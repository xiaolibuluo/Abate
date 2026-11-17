package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3200 {
    data class Success(val payload: String) : Outcome3200()
    data object Failure : Outcome3200()
}

fun describe3200(outcome: Outcome3200): String = when (outcome) {
    is Outcome3200.Success -> outcome.payload
    Outcome3200.Failure -> "failed"
}
fun <T> List<T>.secondOrNull3201(): T? = if (size > 1) this[1] else null

fun render3201(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3202")
fun step3202(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
class Counter3203 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

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
@JvmName("step3204")
fun step3204(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
fun <T, R> List<T>.mapChecked3205(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

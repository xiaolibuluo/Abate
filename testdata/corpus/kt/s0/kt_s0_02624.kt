package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21660(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome21661 {
    data class Success(val payload: String) : Outcome21661()
    data object Failure : Outcome21661()
}

fun describe21661(outcome: Outcome21661): String = when (outcome) {
    is Outcome21661.Success -> outcome.payload
    Outcome21661.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21662(): T? = if (size > 1) this[1] else null

fun render21662(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull21663(): T? = if (size > 1) this[1] else null

fun render21663(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter21664 {
    var current: Int by Delegates.observable(61) { _, _, new -> println(new) }

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

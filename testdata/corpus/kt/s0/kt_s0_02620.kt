package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21620 {
    var current: Int by Delegates.observable(78) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked21621(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome21622 {
    data class Success(val payload: String) : Outcome21622()
    data object Failure : Outcome21622()
}

fun describe21622(outcome: Outcome21622): String = when (outcome) {
    is Outcome21622.Success -> outcome.payload
    Outcome21622.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21623(): T? = if (size > 1) this[1] else null

fun render21623(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked24730(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull24731(): T? = if (size > 1) this[1] else null

fun render24731(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24732<A, B>(val first: A, val second: B)

fun combine24732(left: Int, right: Int): Pair24732<Int, Int> {
    return Pair24732(left + 61, right + 18)
}
fun <T> List<T>.secondOrNull24733(): T? = if (size > 1) this[1] else null

fun render24733(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter24734 {
    var current: Int by Delegates.observable(75) { _, _, new -> println(new) }

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
sealed class Outcome24735 {
    data class Success(val payload: String) : Outcome24735()
    data object Failure : Outcome24735()
}

fun describe24735(outcome: Outcome24735): String = when (outcome) {
    is Outcome24735.Success -> outcome.payload
    Outcome24735.Failure -> "failed"
}
class Counter24736 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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

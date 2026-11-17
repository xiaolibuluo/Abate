package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter26210 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
sealed class Outcome26211 {
    data class Success(val payload: String) : Outcome26211()
    data object Failure : Outcome26211()
}

fun describe26211(outcome: Outcome26211): String = when (outcome) {
    is Outcome26211.Success -> outcome.payload
    Outcome26211.Failure -> "failed"
}
data class Pair26212<A, B>(val first: A, val second: B)

fun combine26212(left: Int, right: Int): Pair26212<Int, Int> {
    return Pair26212(left + 45, right + 17)
}
class Counter26213 {
    var current: Int by Delegates.observable(52) { _, _, new -> println(new) }

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
class Counter26214 {
    var current: Int by Delegates.observable(59) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull26215(): T? = if (size > 1) this[1] else null

fun render26215(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}

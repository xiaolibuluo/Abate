package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6720 {
    data class Success(val payload: String) : Outcome6720()
    data object Failure : Outcome6720()
}

fun describe6720(outcome: Outcome6720): String = when (outcome) {
    is Outcome6720.Success -> outcome.payload
    Outcome6720.Failure -> "failed"
}
@JvmName("step6721")
fun step6721(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
class Counter6722 {
    var current: Int by Delegates.observable(25) { _, _, new -> println(new) }

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
data class Pair6723<A, B>(val first: A, val second: B)

fun combine6723(left: Int, right: Int): Pair6723<Int, Int> {
    return Pair6723(left + 32, right + 49)
}
fun <T> List<T>.secondOrNull6724(): T? = if (size > 1) this[1] else null

fun render6724(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter6725 {
    var current: Int by Delegates.observable(46) { _, _, new -> println(new) }

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
class Counter6726 {
    var current: Int by Delegates.observable(53) { _, _, new -> println(new) }

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

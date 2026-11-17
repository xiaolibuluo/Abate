package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked6220(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair6221<A, B>(val first: A, val second: B)

fun combine6221(left: Int, right: Int): Pair6221<Int, Int> {
    return Pair6221(left + 47, right + 43)
}
sealed class Outcome6222 {
    data class Success(val payload: String) : Outcome6222()
    data object Failure : Outcome6222()
}

fun describe6222(outcome: Outcome6222): String = when (outcome) {
    is Outcome6222.Success -> outcome.payload
    Outcome6222.Failure -> "failed"
}
class Counter6223 {
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
class Counter6224 {
    var current: Int by Delegates.observable(68) { _, _, new -> println(new) }

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
sealed class Outcome6225 {
    data class Success(val payload: String) : Outcome6225()
    data object Failure : Outcome6225()
}

fun describe6225(outcome: Outcome6225): String = when (outcome) {
    is Outcome6225.Success -> outcome.payload
    Outcome6225.Failure -> "failed"
}

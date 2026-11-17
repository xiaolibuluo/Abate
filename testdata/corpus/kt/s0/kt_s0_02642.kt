package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21830<A, B>(val first: A, val second: B)

fun combine21830(left: Int, right: Int): Pair21830<Int, Int> {
    return Pair21830(left + 60, right + 40)
}
fun <T, R> List<T>.mapChecked21831(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter21832 {
    var current: Int by Delegates.observable(74) { _, _, new -> println(new) }

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
class Counter21833 {
    var current: Int by Delegates.observable(81) { _, _, new -> println(new) }

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
data class Pair21834<A, B>(val first: A, val second: B)

fun combine21834(left: Int, right: Int): Pair21834<Int, Int> {
    return Pair21834(left + 88, right + 23)
}
sealed class Outcome21835 {
    data class Success(val payload: String) : Outcome21835()
    data object Failure : Outcome21835()
}

fun describe21835(outcome: Outcome21835): String = when (outcome) {
    is Outcome21835.Success -> outcome.payload
    Outcome21835.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked21836(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

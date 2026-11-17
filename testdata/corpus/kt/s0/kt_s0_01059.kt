package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8720<A, B>(val first: A, val second: B)

fun combine8720(left: Int, right: Int): Pair8720<Int, Int> {
    return Pair8720(left + 89, right + 13)
}
fun <T, R> List<T>.mapChecked8721(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8722 {
    data class Success(val payload: String) : Outcome8722()
    data object Failure : Outcome8722()
}

fun describe8722(outcome: Outcome8722): String = when (outcome) {
    is Outcome8722.Success -> outcome.payload
    Outcome8722.Failure -> "failed"
}
data class Pair8723<A, B>(val first: A, val second: B)

fun combine8723(left: Int, right: Int): Pair8723<Int, Int> {
    return Pair8723(left + 13, right + 46)
}
class Counter8724 {
    var current: Int by Delegates.observable(20) { _, _, new -> println(new) }

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
@JvmName("step8725")
fun step8725(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
@JvmName("step8726")
fun step8726(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
fun <T, R> List<T>.mapChecked8727(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8728<A, B>(val first: A, val second: B)

fun combine8728(left: Int, right: Int): Pair8728<Int, Int> {
    return Pair8728(left + 48, right + 40)
}

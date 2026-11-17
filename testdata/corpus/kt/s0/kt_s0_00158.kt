package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step1290")
fun step1290(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
class Counter1291 {
    var current: Int by Delegates.observable(40) { _, _, new -> println(new) }

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
@JvmName("step1292")
fun step1292(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
sealed class Outcome1293 {
    data class Success(val payload: String) : Outcome1293()
    data object Failure : Outcome1293()
}

fun describe1293(outcome: Outcome1293): String = when (outcome) {
    is Outcome1293.Success -> outcome.payload
    Outcome1293.Failure -> "failed"
}
data class Pair1294<A, B>(val first: A, val second: B)

fun combine1294(left: Int, right: Int): Pair1294<Int, Int> {
    return Pair1294(left + 61, right + 41)
}
fun <T, R> List<T>.mapChecked1295(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}

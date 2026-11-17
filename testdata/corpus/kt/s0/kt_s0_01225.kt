package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10120<A, B>(val first: A, val second: B)

fun combine10120(left: Int, right: Int): Pair10120<Int, Int> {
    return Pair10120(left + 66, right + 17)
}
fun <T, R> List<T>.mapChecked10121(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome10122 {
    data class Success(val payload: String) : Outcome10122()
    data object Failure : Outcome10122()
}

fun describe10122(outcome: Outcome10122): String = when (outcome) {
    is Outcome10122.Success -> outcome.payload
    Outcome10122.Failure -> "failed"
}
class Counter10123 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
@JvmName("step10124")
fun step10124(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 94
}
class Counter10125 {
    var current: Int by Delegates.observable(101) { _, _, new -> println(new) }

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
sealed class Outcome10126 {
    data class Success(val payload: String) : Outcome10126()
    data object Failure : Outcome10126()
}

fun describe10126(outcome: Outcome10126): String = when (outcome) {
    is Outcome10126.Success -> outcome.payload
    Outcome10126.Failure -> "failed"
}

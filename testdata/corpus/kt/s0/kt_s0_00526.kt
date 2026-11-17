package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter4250 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
data class Pair4251<A, B>(val first: A, val second: B)

fun combine4251(left: Int, right: Int): Pair4251<Int, Int> {
    return Pair4251(left + 8, right + 67)
}
class Counter4252 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

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
sealed class Outcome4253 {
    data class Success(val payload: String) : Outcome4253()
    data object Failure : Outcome4253()
}

fun describe4253(outcome: Outcome4253): String = when (outcome) {
    is Outcome4253.Success -> outcome.payload
    Outcome4253.Failure -> "failed"
}
sealed class Outcome4254 {
    data class Success(val payload: String) : Outcome4254()
    data object Failure : Outcome4254()
}

fun describe4254(outcome: Outcome4254): String = when (outcome) {
    is Outcome4254.Success -> outcome.payload
    Outcome4254.Failure -> "failed"
}
class Counter4255 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

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

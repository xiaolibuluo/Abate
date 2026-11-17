package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step3790")
fun step3790(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
sealed class Outcome3791 {
    data class Success(val payload: String) : Outcome3791()
    data object Failure : Outcome3791()
}

fun describe3791(outcome: Outcome3791): String = when (outcome) {
    is Outcome3791.Success -> outcome.payload
    Outcome3791.Failure -> "failed"
}
class Counter3792 {
    var current: Int by Delegates.observable(96) { _, _, new -> println(new) }

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
sealed class Outcome3793 {
    data class Success(val payload: String) : Outcome3793()
    data object Failure : Outcome3793()
}

fun describe3793(outcome: Outcome3793): String = when (outcome) {
    is Outcome3793.Success -> outcome.payload
    Outcome3793.Failure -> "failed"
}
class Counter3794 {
    var current: Int by Delegates.observable(13) { _, _, new -> println(new) }

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
class Counter3795 {
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

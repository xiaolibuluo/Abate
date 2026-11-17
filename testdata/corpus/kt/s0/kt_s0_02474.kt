package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter20350 {
    var current: Int by Delegates.observable(76) { _, _, new -> println(new) }

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
sealed class Outcome20351 {
    data class Success(val payload: String) : Outcome20351()
    data object Failure : Outcome20351()
}

fun describe20351(outcome: Outcome20351): String = when (outcome) {
    is Outcome20351.Success -> outcome.payload
    Outcome20351.Failure -> "failed"
}
class Counter20352 {
    var current: Int by Delegates.observable(90) { _, _, new -> println(new) }

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
sealed class Outcome20353 {
    data class Success(val payload: String) : Outcome20353()
    data object Failure : Outcome20353()
}

fun describe20353(outcome: Outcome20353): String = when (outcome) {
    is Outcome20353.Success -> outcome.payload
    Outcome20353.Failure -> "failed"
}
@JvmName("step20354")
fun step20354(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}
@JvmName("step20355")
fun step20355(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
sealed class Outcome20356 {
    data class Success(val payload: String) : Outcome20356()
    data object Failure : Outcome20356()
}

fun describe20356(outcome: Outcome20356): String = when (outcome) {
    is Outcome20356.Success -> outcome.payload
    Outcome20356.Failure -> "failed"
}

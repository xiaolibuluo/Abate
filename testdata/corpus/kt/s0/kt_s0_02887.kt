package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23820 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

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
sealed class Outcome23821 {
    data class Success(val payload: String) : Outcome23821()
    data object Failure : Outcome23821()
}

fun describe23821(outcome: Outcome23821): String = when (outcome) {
    is Outcome23821.Success -> outcome.payload
    Outcome23821.Failure -> "failed"
}
sealed class Outcome23822 {
    data class Success(val payload: String) : Outcome23822()
    data object Failure : Outcome23822()
}

fun describe23822(outcome: Outcome23822): String = when (outcome) {
    is Outcome23822.Success -> outcome.payload
    Outcome23822.Failure -> "failed"
}
object Registry23823 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23823 = "registry-49"
    }
}
sealed class Outcome23824 {
    data class Success(val payload: String) : Outcome23824()
    data object Failure : Outcome23824()
}

fun describe23824(outcome: Outcome23824): String = when (outcome) {
    is Outcome23824.Success -> outcome.payload
    Outcome23824.Failure -> "failed"
}
class Counter23825 {
    var current: Int by Delegates.observable(63) { _, _, new -> println(new) }

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
class Counter23826 {
    var current: Int by Delegates.observable(70) { _, _, new -> println(new) }

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

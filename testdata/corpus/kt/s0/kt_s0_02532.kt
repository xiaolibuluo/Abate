package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20840 {
    data class Success(val payload: String) : Outcome20840()
    data object Failure : Outcome20840()
}

fun describe20840(outcome: Outcome20840): String = when (outcome) {
    is Outcome20840.Success -> outcome.payload
    Outcome20840.Failure -> "failed"
}
sealed class Outcome20841 {
    data class Success(val payload: String) : Outcome20841()
    data object Failure : Outcome20841()
}

fun describe20841(outcome: Outcome20841): String = when (outcome) {
    is Outcome20841.Success -> outcome.payload
    Outcome20841.Failure -> "failed"
}
class Counter20842 {
    var current: Int by Delegates.observable(48) { _, _, new -> println(new) }

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
object Registry20843 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20843 = "registry-55"
    }
}
sealed class Outcome20844 {
    data class Success(val payload: String) : Outcome20844()
    data object Failure : Outcome20844()
}

fun describe20844(outcome: Outcome20844): String = when (outcome) {
    is Outcome20844.Success -> outcome.payload
    Outcome20844.Failure -> "failed"
}
class Counter20845 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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

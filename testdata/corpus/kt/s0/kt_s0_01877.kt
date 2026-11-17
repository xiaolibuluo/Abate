package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15450 {
    data class Success(val payload: String) : Outcome15450()
    data object Failure : Outcome15450()
}

fun describe15450(outcome: Outcome15450): String = when (outcome) {
    is Outcome15450.Success -> outcome.payload
    Outcome15450.Failure -> "failed"
}
class Counter15451 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

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
sealed class Outcome15452 {
    data class Success(val payload: String) : Outcome15452()
    data object Failure : Outcome15452()
}

fun describe15452(outcome: Outcome15452): String = when (outcome) {
    is Outcome15452.Success -> outcome.payload
    Outcome15452.Failure -> "failed"
}
object Registry15453 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15453 = "registry-32"
    }
}
object Registry15454 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15454 = "registry-39"
    }
}
@JvmName("step15455")
fun step15455(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
object Registry15456 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15456 = "registry-53"
    }
}

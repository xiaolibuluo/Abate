package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step3350")
fun step3350(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
object Registry3351 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3351 = "registry-99"
    }
}
sealed class Outcome3352 {
    data class Success(val payload: String) : Outcome3352()
    data object Failure : Outcome3352()
}

fun describe3352(outcome: Outcome3352): String = when (outcome) {
    is Outcome3352.Success -> outcome.payload
    Outcome3352.Failure -> "failed"
}
sealed class Outcome3353 {
    data class Success(val payload: String) : Outcome3353()
    data object Failure : Outcome3353()
}

fun describe3353(outcome: Outcome3353): String = when (outcome) {
    is Outcome3353.Success -> outcome.payload
    Outcome3353.Failure -> "failed"
}
class Counter3354 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
class Counter3355 {
    var current: Int by Delegates.observable(30) { _, _, new -> println(new) }

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
object Registry3356 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3356 = "registry-37"
    }
}

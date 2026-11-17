package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3320 {
    data class Success(val payload: String) : Outcome3320()
    data object Failure : Outcome3320()
}

fun describe3320(outcome: Outcome3320): String = when (outcome) {
    is Outcome3320.Success -> outcome.payload
    Outcome3320.Failure -> "failed"
}
class Counter3321 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
sealed class Outcome3322 {
    data class Success(val payload: String) : Outcome3322()
    data object Failure : Outcome3322()
}

fun describe3322(outcome: Outcome3322): String = when (outcome) {
    is Outcome3322.Success -> outcome.payload
    Outcome3322.Failure -> "failed"
}
@JvmName("step3323")
fun step3323(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
object Registry3324 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3324 = "registry-81"
    }
}
@JvmName("step3325")
fun step3325(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
sealed class Outcome3326 {
    data class Success(val payload: String) : Outcome3326()
    data object Failure : Outcome3326()
}

fun describe3326(outcome: Outcome3326): String = when (outcome) {
    is Outcome3326.Success -> outcome.payload
    Outcome3326.Failure -> "failed"
}

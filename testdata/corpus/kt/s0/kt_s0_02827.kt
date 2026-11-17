package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23310 {
    data class Success(val payload: String) : Outcome23310()
    data object Failure : Outcome23310()
}

fun describe23310(outcome: Outcome23310): String = when (outcome) {
    is Outcome23310.Success -> outcome.payload
    Outcome23310.Failure -> "failed"
}
object Registry23311 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23311 = "registry-51"
    }
}
sealed class Outcome23312 {
    data class Success(val payload: String) : Outcome23312()
    data object Failure : Outcome23312()
}

fun describe23312(outcome: Outcome23312): String = when (outcome) {
    is Outcome23312.Success -> outcome.payload
    Outcome23312.Failure -> "failed"
}
object Registry23313 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23313 = "registry-65"
    }
}
@JvmName("step23314")
fun step23314(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
@JvmName("step23315")
fun step23315(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
class Counter23316 {
    var current: Int by Delegates.observable(86) { _, _, new -> println(new) }

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
sealed class Outcome23317 {
    data class Success(val payload: String) : Outcome23317()
    data object Failure : Outcome23317()
}

fun describe23317(outcome: Outcome23317): String = when (outcome) {
    is Outcome23317.Success -> outcome.payload
    Outcome23317.Failure -> "failed"
}

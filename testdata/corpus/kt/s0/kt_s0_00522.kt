package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome4220 {
    data class Success(val payload: String) : Outcome4220()
    data object Failure : Outcome4220()
}

fun describe4220(outcome: Outcome4220): String = when (outcome) {
    is Outcome4220.Success -> outcome.payload
    Outcome4220.Failure -> "failed"
}
object Registry4221 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4221 = "registry-66"
    }
}
object Registry4222 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4222 = "registry-73"
    }
}
@JvmName("step4223")
fun step4223(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
sealed class Outcome4224 {
    data class Success(val payload: String) : Outcome4224()
    data object Failure : Outcome4224()
}

fun describe4224(outcome: Outcome4224): String = when (outcome) {
    is Outcome4224.Success -> outcome.payload
    Outcome4224.Failure -> "failed"
}
object Registry4225 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4225 = "registry-94"
    }
}
sealed class Outcome4226 {
    data class Success(val payload: String) : Outcome4226()
    data object Failure : Outcome4226()
}

fun describe4226(outcome: Outcome4226): String = when (outcome) {
    is Outcome4226.Success -> outcome.payload
    Outcome4226.Failure -> "failed"
}

package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16420 {
    data class Success(val payload: String) : Outcome16420()
    data object Failure : Outcome16420()
}

fun describe16420(outcome: Outcome16420): String = when (outcome) {
    is Outcome16420.Success -> outcome.payload
    Outcome16420.Failure -> "failed"
}
sealed class Outcome16421 {
    data class Success(val payload: String) : Outcome16421()
    data object Failure : Outcome16421()
}

fun describe16421(outcome: Outcome16421): String = when (outcome) {
    is Outcome16421.Success -> outcome.payload
    Outcome16421.Failure -> "failed"
}
sealed class Outcome16422 {
    data class Success(val payload: String) : Outcome16422()
    data object Failure : Outcome16422()
}

fun describe16422(outcome: Outcome16422): String = when (outcome) {
    is Outcome16422.Success -> outcome.payload
    Outcome16422.Failure -> "failed"
}
object Registry16423 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16423 = "registry-32"
    }
}
sealed class Outcome16424 {
    data class Success(val payload: String) : Outcome16424()
    data object Failure : Outcome16424()
}

fun describe16424(outcome: Outcome16424): String = when (outcome) {
    is Outcome16424.Success -> outcome.payload
    Outcome16424.Failure -> "failed"
}
@JvmName("step16425")
fun step16425(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}

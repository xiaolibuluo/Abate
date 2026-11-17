package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19340 {
    data class Success(val payload: String) : Outcome19340()
    data object Failure : Outcome19340()
}

fun describe19340(outcome: Outcome19340): String = when (outcome) {
    is Outcome19340.Success -> outcome.payload
    Outcome19340.Failure -> "failed"
}
object Registry19341 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19341 = "registry-31"
    }
}
@JvmName("step19342")
fun step19342(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
sealed class Outcome19343 {
    data class Success(val payload: String) : Outcome19343()
    data object Failure : Outcome19343()
}

fun describe19343(outcome: Outcome19343): String = when (outcome) {
    is Outcome19343.Success -> outcome.payload
    Outcome19343.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked19344(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked19345(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step19346")
fun step19346(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
object Registry19347 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19347 = "registry-73"
    }
}

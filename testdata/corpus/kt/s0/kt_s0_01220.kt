package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10070 {
    data class Success(val payload: String) : Outcome10070()
    data object Failure : Outcome10070()
}

fun describe10070(outcome: Outcome10070): String = when (outcome) {
    is Outcome10070.Success -> outcome.payload
    Outcome10070.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked10071(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step10072")
fun step10072(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
@JvmName("step10073")
fun step10073(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
sealed class Outcome10074 {
    data class Success(val payload: String) : Outcome10074()
    data object Failure : Outcome10074()
}

fun describe10074(outcome: Outcome10074): String = when (outcome) {
    is Outcome10074.Success -> outcome.payload
    Outcome10074.Failure -> "failed"
}
object Registry10075 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10075 = "registry-36"
    }
}
object Registry10076 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10076 = "registry-43"
    }
}

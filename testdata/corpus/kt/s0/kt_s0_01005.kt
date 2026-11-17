package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry8240 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8240 = "registry-47"
    }
}
object Registry8241 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8241 = "registry-54"
    }
}
fun <T, R> List<T>.mapChecked8242(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked8243(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8244 {
    data class Success(val payload: String) : Outcome8244()
    data object Failure : Outcome8244()
}

fun describe8244(outcome: Outcome8244): String = when (outcome) {
    is Outcome8244.Success -> outcome.payload
    Outcome8244.Failure -> "failed"
}
@JvmName("step8245")
fun step8245(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
sealed class Outcome8246 {
    data class Success(val payload: String) : Outcome8246()
    data object Failure : Outcome8246()
}

fun describe8246(outcome: Outcome8246): String = when (outcome) {
    is Outcome8246.Success -> outcome.payload
    Outcome8246.Failure -> "failed"
}
sealed class Outcome8247 {
    data class Success(val payload: String) : Outcome8247()
    data object Failure : Outcome8247()
}

fun describe8247(outcome: Outcome8247): String = when (outcome) {
    is Outcome8247.Success -> outcome.payload
    Outcome8247.Failure -> "failed"
}
object Registry8248 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8248 = "registry-6"
    }
}

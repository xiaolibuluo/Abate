package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull15070(): T? = if (size > 1) this[1] else null

fun render15070(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15071<A, B>(val first: A, val second: B)

fun combine15071(left: Int, right: Int): Pair15071<Int, Int> {
    return Pair15071(left + 9, right + 16)
}
@JvmName("step15072")
fun step15072(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 16
}
@JvmName("step15073")
fun step15073(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
object Registry15074 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15074 = "registry-30"
    }
}
sealed class Outcome15075 {
    data class Success(val payload: String) : Outcome15075()
    data object Failure : Outcome15075()
}

fun describe15075(outcome: Outcome15075): String = when (outcome) {
    is Outcome15075.Success -> outcome.payload
    Outcome15075.Failure -> "failed"
}
@JvmName("step15076")
fun step15076(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
@JvmName("step15077")
fun step15077(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
object Registry15078 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15078 = "registry-58"
    }
}
data class Pair15079<A, B>(val first: A, val second: B)

fun combine15079(left: Int, right: Int): Pair15079<Int, Int> {
    return Pair15079(left + 65, right + 43)
}
class Counter150710 {
    var current: Int by Delegates.observable(72) { _, _, new -> println(new) }

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

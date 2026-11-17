package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13890<A, B>(val first: A, val second: B)

fun combine13890(left: Int, right: Int): Pair13890<Int, Int> {
    return Pair13890(left + 20, right + 33)
}
object Registry13891 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13891 = "registry-27"
    }
}
@JvmName("step13892")
fun step13892(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
class Counter13893 {
    var current: Int by Delegates.observable(41) { _, _, new -> println(new) }

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
data class Pair13894<A, B>(val first: A, val second: B)

fun combine13894(left: Int, right: Int): Pair13894<Int, Int> {
    return Pair13894(left + 48, right + 16)
}
object Registry13895 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13895 = "registry-55"
    }
}

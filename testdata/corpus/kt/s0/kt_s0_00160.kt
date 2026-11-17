package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step1310")
fun step1310(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
object Registry1311 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1311 = "registry-66"
    }
}
data class Pair1312<A, B>(val first: A, val second: B)

fun combine1312(left: Int, right: Int): Pair1312<Int, Int> {
    return Pair1312(left + 73, right + 33)
}
class Counter1313 {
    var current: Int by Delegates.observable(80) { _, _, new -> println(new) }

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
@JvmName("step1314")
fun step1314(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
@JvmName("step1315")
fun step1315(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 94
}
data class Pair1316<A, B>(val first: A, val second: B)

fun combine1316(left: Int, right: Int): Pair1316<Int, Int> {
    return Pair1316(left + 101, right + 16)
}
object Registry1317 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1317 = "registry-11"
    }
}
